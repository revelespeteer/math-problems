#include <iostream>
using namespace std;

int main() {
    int n;
    cout << "Enter an integer: ";
    cin >> n;

    if (n % 2 == 0) {
        for (int i = 1; i <= n; i += 2) {
            if (i % 3 == 0 || i % 5 == 0) {
                cout << "Odd multiples of both 3 and 5:" << endl;
                break;
            }
            else {
                cout << "Even numbers: ";
                for (int j = 1; j <= n; j += 2) {
                    if (j % 3 == 0 || j % 5 == 0) {
                        cout << j << " ";
                    }
                }
                break;
            }
        }
    }
    else {
        cout << "Odd numbers: ";
        for (int i = 1; i <= n; i += 2) {
            if (i % 3 == 0 || i % 5 == 0) {
                cout << i << " ";
            }
        }
        cout << endl;
    }

    return 0;
}
