import { getRandomInteger } from "./utils";

export type KnapsackInstance = {
    values: number[];
    weights: number[];
    capacity: number;
};

export function generateKnapsackInstances(instances: number): KnapsackInstance[] {
    const result: KnapsackInstance[] = [];

    for (let i = 0; i < instances; i++) {
        const n = getRandomInteger(5, 50);
        let totalWeight = 0;

        const values: number[] = [];
        const weights: number[] = [];

        for (let j = 0; j < n; j++) {
            const value = getRandomInteger(1, 100);
            const weight = getRandomInteger(1, 50);
            values.push(value);
            weights.push(weight);
            totalWeight += weight;
        }

        const ratio = Math.random() * (0.7 - 0.3) + 0.3;
        const capacity = Math.floor(totalWeight * ratio);

        result.push({
            values: values,
            weights: weights,
            capacity: capacity
        });
    }

    return result;
}