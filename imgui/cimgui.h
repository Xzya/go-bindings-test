#define API

#if defined __cplusplus
#define EXTERN extern "C"
#else
#define EXTERN extern
#endif

#define CIMGUI_API EXTERN API

struct ImVec2;

#ifdef CIMGUI_DEFINE_ENUMS_AND_STRUCTS
struct ImVec2
{
  float x, y;
};
#endif

CIMGUI_API void foo(struct ImVec2 *out);