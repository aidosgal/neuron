#include "clips.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

char* GetFactSlotString(Fact* fact, const char* slotName) {
    CLIPSValue cv;
    char buffer[1024];
    const char* sourceStr = NULL;
    
    if (fact == NULL || slotName == NULL) {
        return NULL;
    }
    
    if (GetFactSlot(fact, slotName, &cv) != GSE_NO_ERROR) {
        return NULL;
    }
    
    if (CVIsType(&cv, SYMBOL_BIT) || CVIsType(&cv, STRING_BIT) || CVIsType(&cv, INSTANCE_NAME_BIT)) {
        if (cv.lexemeValue != NULL && cv.lexemeValue->contents != NULL) {
            sourceStr = cv.lexemeValue->contents;
        }
    }
    else if (CVIsType(&cv, INTEGER_BIT)) {
        if (cv.integerValue != NULL) {
            snprintf(buffer, sizeof(buffer), "%lld", cv.integerValue->contents);
            sourceStr = buffer;
        }
    }
    else if (CVIsType(&cv, FLOAT_BIT)) {
        if (cv.floatValue != NULL) {
            snprintf(buffer, sizeof(buffer), "%f", cv.floatValue->contents);
            sourceStr = buffer;
        }
    }
    
    if (sourceStr != NULL) {
        char* result = (char*)malloc(strlen(sourceStr) + 1);
        if (result != NULL) {
            strcpy(result, sourceStr);
        }
        return result;
    }
    
    return NULL;
}

void FreeString(char* str) {
    if (str != NULL) {
        free(str);
    }
}

Fact* GetNextFactWrapper(Environment* env, Fact* fact) {
    return GetNextFact(env, fact);
}

Deftemplate* GetFactDeftemplate(Fact* fact) {
    return FactDeftemplate(fact);
}

const char* GetDeftemplateName(Deftemplate* deftemplate) {
    return DeftemplateName(deftemplate);
}
