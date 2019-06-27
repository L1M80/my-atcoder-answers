#include <string.h>
#include <iostream>
#include <vector>
#define N_MAX_DIGIT 5

using namespace std;

void apart(vector<int> &aparted, int toBeAparted) {
    for (int i = 0; i < aparted.size(); i++) {
        aparted[i] = toBeAparted % 10;
        toBeAparted /= 10;
    }
}

int sum(vector<int> toBeSummed) {
    int summed = 0;
    for (int i = 0; i < toBeSummed.size(); i++) {
        summed += toBeSummed[i];
    }
    return summed;
}

int main() {
    int n = 0;
    int a = 0;
    int b = 0;
    int result;
    vector<int> aparted(N_MAX_DIGIT, 0);

    scanf("%d%d%d", &n, &a, &b);

    for (int i = 1; i <= n; i++) {
        apart(aparted, i);
        int summed = sum(aparted);
        if (summed >= a && summed <= b) {
            result += i;
        }
    }

    printf("%d", result);

    return 0;
}
