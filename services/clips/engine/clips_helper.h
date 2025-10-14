#ifndef CLIPS_HELPERS_H
#define CLIPS_HELPERS_H

#include "clips.h"

char* GetFactSlotString(Fact* fact, const char* slotName);
void FreeString(char* str);
Fact* GetNextFactWrapper(Environment* env, Fact* fact);
Deftemplate* GetFactDeftemplate(Fact* fact);
const char* GetDeftemplateName(Deftemplate* deftemplate);

#endif
