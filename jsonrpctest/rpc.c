#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include "rpc.h"

int main() {
    GoString action = {"condenser_api.get_accounts", strlen("condenser_api.get_accounts") + 1};

    GoString arg1 = {"[\"wb-100\", \"wb-200\"]", strlen("[\"wb-100\", \"wb-200\"]") + 1};
    GoString data[1] = {arg1};
    GoSlice args = {data, 1};

    GoString url = {"ws://8.8.8.8:38090", strlen("ws://8.8.8.8:38090") + 1};

    char* result = call(url, action, args); 
    printf("%s\n", result);

    free(result);
} 