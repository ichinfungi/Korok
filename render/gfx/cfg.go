package gfx

const (
	CONFIG_RENDERER_OPENGLES = 20
)

const CONFIG_MAX_DRAW_CALLS 				uint32 = 1 << 10

const (
	CONFIG_MAX_COMMAND_BUFFER_SIZE uint32 = 64 << 10

	CONFIG_DYNAMIC_INDEX_BUFFER_SIZE 	uint32 = 1 << 20
	CONFIG_DYNAMIC_VERTEX_BUFFER_SIZE 	uint32 = 1 << 20
	CONFIG_TRANSIENT_INDEX_BUFFER_SIZE 	uint32 = 2 << 10
	CONFIG_TRANSIENT_VERTEX_BUFFER_SIZE uint32 = 2 << 10
)

const (
	CONFIG_MAX_INDEX_BUFFERS 			uint32 = 1 << 10
	CONFIG_MAX_VERTEX_BUFFERS 			uint32 = 1 << 10
	CONFIG_MAX_DYNAMIC_INDEX_BUFFERS 	uint32 = 1 << 10
	CONFIG_MAX_DYNAMIC_VERTEX_BUFFERS 	uint32 = 1 << 10

	CONFIG_MAX_VIEWS    uint32 = 32
	CONFIG_MAX_SAMPLERS uint32 = 4
	CONFIG_MAX_STREAMS  uint32 = 2

	CONFIG_MAX_UNIFORMS 	uint32 = 512
	CONFIG_MAX_SHADERS 		uint32 = 32
	CONFIG_MAX_PROGRAMS 	uint32 = 32
	CONFIG_MAX_TEXTURES 	uint32 = 1 << 10

	CONFIG_MAX_VERTEX_LAYOUT  uint32 = 32
	CONFIG_MAX_COLOR_PALETTE  uint32 = 16
	CONFIG_MAX_FRAME_BUFFERS  uint32 = 32
	CONFIG_MAX_FB_ATTACHMENTS uint32 = 8
)