/**
 * @prettier
 * @flow
 */

const RED = 1;
const BLACK = 2;
const BOARD_WIDTH = 7;
const BOARD_HEIGHT = 6;

const stringIntMap = {
    "0": 0,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
};

// type Board struct {
//     Data []int
// }
class Board {
    data: Array<number>;
    constructor(data: ?Array<number>) {
        this.data = data || [];
    }

    playMove(column: number): ?Error {
        if (this.data.length == BOARD_WIDTH * BOARD_HEIGHT) {
            return new Error("Board full");
        }
        let count = 0;
        for (var i = 0; i < this.data.length; i++) {
            if (this.data[i] == column) {
                count++;
            }
        }
        if (count == BOARD_HEIGHT) {
            return new Error("Column full");
        }
        this.data.push(column);
    }
    whoIsNext(): number {
        if (this.data.length % 2 == 1) {
            return BLACK;
        }
        return RED;
    }

    asMatrix(): Array<Array<number>> {
        let counters = new Array(BOARD_WIDTH).fill(0);
        let board = new Array(BOARD_WIDTH).fill(
            new Array(BOARD_HEIGHT).fill(0),
        );
        for (var i = 0; i < this.data.length; i++) {
            const column = this.data[i];
            board[column][counters[column]] = i % 2 + 1;
            counters[column] += 1;
        }
        return board;
    }
    isWon(): number {
        const board = this.asMatrix();
        // horizontalCheck
        for (let j = 0; j < BOARD_HEIGHT - 3; j++) {
            for (let i = 0; i < BOARD_WIDTH; i++) {
                if (
                    (board[i][j] == RED || board[i][j] == BLACK) &&
                    board[i][j] == board[i][j + 1] &&
                    board[i][j] == board[i][j + 2] &&
                    board[i][j] == board[i][j + 3]
                ) {
                    return board[i][j];
                }
            }
        }
        // verticalCheck
        for (let i = 0; i < BOARD_WIDTH - 3; i++) {
            for (let j = 0; j < BOARD_HEIGHT; j++) {
                if (
                    (board[i][j] == RED || board[i][j] == BLACK) &&
                    board[i][j] == board[i + 1][j] &&
                    board[i][j] == board[i + 2][j] &&
                    board[i][j] == board[i + 3][j]
                ) {
                    return board[i][j];
                }
            }
        }
        // ascendingDiagonalCheck
        for (let i = 3; i < BOARD_WIDTH; i++) {
            for (let j = 0; j < BOARD_HEIGHT - 3; j++) {
                if (
                    (board[i][j] == RED || board[i][j] == BLACK) &&
                    board[i][j] == board[i - 1][j + 1] &&
                    board[i][j] == board[i - 2][j + 2] &&
                    board[i][j] == board[i - 3][j + 3]
                ) {
                    return board[i][j];
                }
            }
        }
        // descendingDiagonalCheck
        for (let i = 3; i < BOARD_WIDTH; i++) {
            for (let j = 3; j < BOARD_HEIGHT; j++) {
                if (
                    (board[i][j] == RED || board[i][j] == BLACK) &&
                    board[i][j] == board[i - 1][j - 1] &&
                    board[i][j] == board[i - 2][j - 2] &&
                    board[i][j] == board[i - 3][j - 3]
                ) {
                    return board[i][j];
                }
            }
        }
        return 0;
    }
}

const BoardFromString = (moves: string): Board => {
    let board = new Board();
    for (var i = 0; i < moves.length; i++) {
        board.data.push(stringIntMap[moves]);
    }
    return board;
};

module.exports = {
    Board: Board,
    BoardFromString: BoardFromString,
};
