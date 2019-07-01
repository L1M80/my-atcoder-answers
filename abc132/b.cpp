#include <iostream>
#include <vector>
using namespace std;

int main() {
    int answer = 0;
    int n;
    cin >> n;
    vector<int> p(n);
    for (int i = 0; i < n; i++) cin >> p[i];

    for (int i = 1; i < n - 1; i++) {
        if ((p[i] < p[i + 1] && p[i - 1] < p[i]) ||
            (p[i] > p[i + 1] && p[i - 1] > p[i]))
            answer += 1;
    }

    cout << answer << endl;
}
