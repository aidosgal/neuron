#include <stdio.h>
#include <stdlib.h>
#include "clips.h"
#include "cJSON.h"

int main(void)
{
    void *env = CreateEnvironment();
    if (env == NULL)
    {
        fprintf(stderr, "failed to create CLIPS environment\n");
        return 1;
    }

    cJSON *root = cJSON_CreateObject();
    cJSON_AddStringToObject(root, "status", "running");
    char *jsonStr = cJSON_Print(root);
    printf("JSON Test: %s\n", jsonStr);

    if (!Load(env, "rules/traffic.clp"))
    {
        fprintf(stderr, "failed to load rules file.\n");
        DestroyEnvironment(env);
        return 1;
    }

    Reset(env);
    Run(env, -1);

    cJSON_Delete(root);
    free(jsonStr);
    DestroyEnvironment(env);

    printf("done.\n");
    return 0;
}
