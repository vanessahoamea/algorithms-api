import http from "k6/http";
import { SharedArray } from "k6/data";
import { check, group, sleep } from "k6";
import { getRandomInteger } from "./helpers/utils";
import { generateNQueensInstances } from "./helpers/n_queens";
import { generateKnapsackInstances } from "./helpers/knapsack";
import { generateShortestPathInstances } from "./helpers/shortest_path";

const BASE_URL = __ENV.BASE_URL;
const OPTIONS_FILE = `../configs/${__ENV.OPTIONS_FILE}`;
const DEFAULT_HEADERS = {
    "Content-Type": "application/json"
};

export const options = JSON.parse(open(OPTIONS_FILE));

const nQueensInstances = new SharedArray("N Queens problem instances", () => {
    return generateNQueensInstances(100);
});
const knapsackInstances = new SharedArray("Knapsack problem instances", () => {
    return generateKnapsackInstances(100);
});
const shortestPathInstances = new SharedArray("Shortest Path problem instances", () => {
    return generateShortestPathInstances(100);
});

export default function() {
    group("N Queens endpoint", () => {
        const problemIndex = getRandomInteger(0, nQueensInstances.length - 1);
        const body = JSON.stringify(nQueensInstances[problemIndex]);
        const response = http.post(`${BASE_URL}/n-queens`, body, { headers: DEFAULT_HEADERS });
        check(response, {
            "status code is 200": (r) => r.status === 200
        });
    });

    group("Knapsack endpoint", () => {
        const problemIndex = getRandomInteger(0, knapsackInstances.length - 1);
        const body = JSON.stringify(knapsackInstances[problemIndex]);
        const response = http.post(`${BASE_URL}/knapsack`, body, { headers: DEFAULT_HEADERS });
        check(response, {
            "status code is 200": (r) => r.status === 200
        });
    });

    group("Shortest Path endpoint", () => {
        const problemIndex = getRandomInteger(0, shortestPathInstances.length - 1);
        const body = JSON.stringify(shortestPathInstances[problemIndex]);
        const response = http.post(`${BASE_URL}/shortest-path`, body, { headers: DEFAULT_HEADERS });
        check(response, {
            "status code is 200": (r) => r.status === 200
        });
    });

    // slight delay until the next user becomes active
    const duration = Math.random();
    sleep(duration);
}