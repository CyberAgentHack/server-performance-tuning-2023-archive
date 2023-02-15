import http from 'k6/http';
import {check, fail} from 'k6';
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';

export const options = {
  thresholds: {
      http_req_duration: ['p(95)<1000'], // 95% のリクエストは 1000ms (1s) 以内に収める
  },
  scenarios: {
      contacts: {
          executor: 'ramping-arrival-rate', // https://k6.io/docs/using-k6/scenarios/executors/ramping-arrival-rate
          exec: 'load_test',

          gracefulStop: '10s',

          preAllocatedVUs: 20,
          stages: [
              // target: 1 秒あたりの load_test 関数の実行回数の目標値
              // duration: target 到達までにかかる時間
              { target: 5, duration: '1m' }, // 1分かけてload_test関数の実行回数を5まで大きくする
              { target: 5, duration: '1m' }, // 1秒あたりの実行回数5回を1分間維持する 
          ],
      },
  },
};


export function load_test() {
  const seriesURL =  new URL(`${__ENV.API_BASE_URL}/series`);
  const offset = Math.floor(Math.random() * 2183)
  seriesURL.searchParams.append(`limit`, `20`);
  seriesURL.searchParams.append(`offset`, `${offset}`);
  
  let res = http.get(seriesURL.toString());

  check_status_ok(res);
  
  let body = res.json();
  
  const seasonRequests = Array();
  body.series.forEach(series => {
    const url = new URL(`${__ENV.API_BASE_URL}/seasons`);
    url.searchParams.append(`limit`, `20`)
    url.searchParams.append(`offset`, `0`)
    url.searchParams.append(`seriesId`, `${series.id}`)
    seasonRequests.push(['GET', url.toString()])
  });
  
  let seasonRes = http.batch(seasonRequests);
  seasonRes.forEach(
    (res) => check_status_ok(res)
  )
  
  const episodeRequests = Array();
  seasonRes.forEach(
    (res) => {
      const body = res.json();
      if (body === null || body.seasons === null) {
        return;
      }
      body.seasons.forEach((season) => {
        const url = new URL(`${__ENV.API_BASE_URL}/episodes`);
        url.searchParams.append(`limit`, `20`)
        url.searchParams.append(`offset`, `0`)
        url.searchParams.append(`seasonId`, `${season.id}`)
        episodeRequests.push(['GET', url.toString()])
      })
    }
  )
  
  let episodeRes = http.batch(episodeRequests);
  episodeRes.forEach(
    (res) => check_status_ok(res)
  )
}

function check_status_ok(res) {
  const result = check(res, {
    'is status OK': (r) => r.status === 200,
  }); 
  if (!result) {
    fail(`status is not 200. response: ${JSON.stringify(res)}`);
  }
}
