import http from 'k6/http';
import { sleep } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";

export const options = {
  stages: [
    { duration: '0.2m', target: 100 }, // below normal load
    { duration: '0.2m', target: 200 }, // normal load
    { duration: '0.2m', target: 300 }, // around the breaking point
    { duration: '0.2m', target: 400 }, // beyond the breaking point
    { duration: '0.2m', target: 400 }, // beyond the breaking point
    { duration: '0.2m', target: 0 }, // scale down. Recovery stage.
  ],
};

export default function () {
    http.post('http://localhost:8080/ping')
    sleep(1);
}

export function handleSummary(data) {
    return {
        "result.html": htmlReport(data),
        stdout: textSummary(data, { indent: " ", enableColors: true }),
      };
    
}
  
