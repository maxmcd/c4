/**
 * @prettier
 * @flow
 */

import board from "./board";

test("BoardIsWon", () => {
    const tests = [
        "0011223",
        "0101010",
        "333333215444442122115620552",
        "444444326555553233226131666",
    ];
    for (let i = 0; i < tests.length; i++) {
        let test = tests[i];
        let b = board.BoardFromString(test);
        if (b.isWon() !== board.RED) {
            throw new Error("Should be a winner " + test);
        }
    }
});

