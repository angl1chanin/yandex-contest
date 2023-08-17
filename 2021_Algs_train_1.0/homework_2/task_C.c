#include <stdio.h>
#include <stdlib.h>

int main() {
    int n,x;
    scanf("%d",&n);
    int *arr = malloc(sizeof(int)*n);
    for (int i=0; i<n; i++) {
        scanf("%d",&arr[i]);
    }
    scanf("%d",&x);

    int *min_diffs = malloc(sizeof(int)*n);
    for (int i=0; i<n; i++) {
        min_diffs[i] = abs(arr[i] - x);
    }

    int min_diff = min_diffs[0];
    int index = 0;
    for (int i=0; i<n; i++) {
        if (min_diffs[i] < min_diff) {
            min_diff = min_diffs[i];
            index = i;
        }
    }
    printf("%d", arr[index]);
    free(min_diffs);
    free(arr);
    return 0;
}