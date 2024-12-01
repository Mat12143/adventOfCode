#include <iostream>
#include <fstream>
using namespace std;

#define LEN 1000

void bubbleSort(int v[LEN]) {
    for (int i = 0; i < LEN - 1; i++) {
        for (int j = 0; j < LEN - i - 1; j++) {
            if (v[j] > v[j + 1]) swap(v[j], v[j + 1]);
        }
    }
}

int main(int argc, char** argv) {

    fstream input;

    input.open("long.txt", ios::in);
    if (input.fail()) {
        cerr << "Error while reading the file" << endl;
        exit(1);
    }

    int left[LEN] = {0}, right[LEN] = {0};
    int lItem, rItem;
    int differencies = 0;

    for(int i = 0; i < LEN; i++) {
        input >> lItem >> rItem;

        left[i] = lItem;
        right[i] = rItem;
    }

    // Sort arrays
    bubbleSort(left);
    bubbleSort(right);

    for(int i = 0; i < LEN; i++) {
        int diff = left[i] - right[i];

        if (diff < 0) diff *= -1;

        differencies += diff;

    }

    cout << "DIFF: " << differencies << endl;

    input.close();

    return 0;
}
