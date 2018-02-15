#include "imgui.h"
#include "cimgui.h"

CIMGUI_API void foo(ImVec2 *out) {
  *out = ImGui::foo();
}
