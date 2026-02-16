import { getRandomInteger } from "./utils";

export type ShortestPathInstance = {
    n: number;
    edges: number[][];
    souce: number;
};

export function generateShortestPathInstances(instances: number): ShortestPathInstance[] {
    const result: ShortestPathInstance[] = [];

    for (let i = 0; i < instances; i++) {
        const n = getRandomInteger(6, 50);
        const source = getRandomInteger(0, n - 1);
        const edges: number[][] = [];

        const edgeCount = getRandomInteger(1, n * (n - 1));
        const edgeSet = new Set<string>();
        while (edgeSet.size < edgeCount) {
            const from = getRandomInteger(0, n - 1);
            const to = getRandomInteger(0, n - 1);

            if (from !== to) {
                const key = `${from},${to}`;
                if (!edgeSet.has(key)) {
                    const weight = getRandomInteger(1, 100);
                    edges.push([from, to, weight]);
                    edgeSet.add(key);
                }
            }
        }

        result.push({
            n: n,
            edges: edges,
            souce: source
        });
    }

    return result;
}