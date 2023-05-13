const description = "Sums up value1 and value2.";

const options = [
    { type: TYPES.Number, name: "value1", description: "The first value to sum.", required: true },
    { type: TYPES.Number, name: "value2", description: "The second value to sum.", required: true }
]

const execute = (args) => {
    const value1 = args.value1;
    const value2 = args.value2;

    const sum = value1 + value2
    return "" + value1 + " + " + value2 + " = " + sum
}