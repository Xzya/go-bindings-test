#pragma once

#ifndef IMGUI_API
#define IMGUI_API
#endif

struct ImVec2
{
  float x, y;
  ImVec2() { x = y = 0.0f; }
  ImVec2(float _x, float _y) { x = _x; y = _y; }
};

namespace ImGui
{
  IMGUI_API ImVec2 foo();
}