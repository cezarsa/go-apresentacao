#include <stdio.h>
#include <stdlib.h>

typedef struct{int a;} myStruct;

myStruct fn() {
    myStruct m;
    return m;
}

int main() {
    FILE*(*myOpener)(const char*, const char*);
    myOpener = fopen;

    myStruct m = (myStruct){5};
    FILE *file = myOpener("arquivo.txt", "rb");
    fseek(file, 0, SEEK_END);
    long size = ftell(file);
    rewind(file);
    char *buffer = (char*)malloc(size+1);
    fread(buffer, size, 1, file);
    buffer[size] = 0;
    printf("%s\n", buffer);
    free(buffer);
    return 0;
}