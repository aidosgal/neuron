#include "clips.h"
#include <stdlib.h>

const char* GetFactSlotString(Fact* fact, const char* slotName) {
    CLIPSValue cv;
    if (GetFactSlot(fact, slotName, &cv) == GSE_NO_ERROR) {
        if (cv.lexemeValue != NULL) {
            return cv.lexemeValue->contents;
        }
    }
    return NULL;
}
