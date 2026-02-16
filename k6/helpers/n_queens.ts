import { getRandomInteger } from "./utils";

export type NQueensInstance = {
    n: number;
    blocked: number[][];
};

export function generateNQueensInstances(instances: number): NQueensInstance[] {
    const result: NQueensInstance[] = [];

    for (let i = 0; i < instances; i++) {
        const n = getRandomInteger(4, 50);
        const blockedCount = getRandomInteger(1, (n * n) / 2);

        const blocked = new Set<string>();
        while (blocked.size < blockedCount) {
            const row = getRandomInteger(0, n - 1);
            const col = getRandomInteger(0, n - 1);
            blocked.add(`${row},${col}`);
        }

        result.push({
            n: n,
            blocked: Array.from(blocked).map(coord => coord.split(',').map(Number))
        });
    }

    return result;
}