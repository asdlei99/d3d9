package d3d9

//#include <d3d9.h>
import "C"

const (
	ADAPTER_DEFAULT                             = C.D3DADAPTER_DEFAULT
	CAPS_READ_SCANLINE                          = C.D3DCAPS_READ_SCANLINE
	CAPS2_CANAUTOGENMIPMAP                      = C.D3DCAPS2_CANAUTOGENMIPMAP
	CAPS2_CANCALIBRATEGAMMA                     = C.D3DCAPS2_CANCALIBRATEGAMMA
	CAPS2_CANMANAGERESOURCE                     = C.D3DCAPS2_CANMANAGERESOURCE
	CAPS2_DYNAMICTEXTURES                       = C.D3DCAPS2_DYNAMICTEXTURES
	CAPS2_FULLSCREENGAMMA                       = C.D3DCAPS2_FULLSCREENGAMMA
	CAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD      = C.D3DCAPS3_ALPHA_FULLSCREEN_FLIP_OR_DISCARD
	CAPS3_COPY_TO_VIDMEM                        = C.D3DCAPS3_COPY_TO_VIDMEM
	CAPS3_COPY_TO_SYSTEMMEM                     = C.D3DCAPS3_COPY_TO_SYSTEMMEM
	CAPS3_LINEAR_TO_SRGB_PRESENTATION           = C.D3DCAPS3_LINEAR_TO_SRGB_PRESENTATION
	CLEAR_STENCIL                               = C.D3DCLEAR_STENCIL
	CLEAR_TARGET                                = C.D3DCLEAR_TARGET
	CLEAR_ZBUFFER                               = C.D3DCLEAR_ZBUFFER
	CREATE_ADAPTERGROUP_DEVICE                  = 0x00000200
	CREATE_DISABLE_DRIVER_MANAGEMENT            = 0x00000100
	CREATE_DISABLE_DRIVER_MANAGEMENT_EX         = 0x00000400
	CREATE_FPU_PRESERVE                         = C.D3DCREATE_FPU_PRESERVE
	CREATE_HARDWARE_VERTEXPROCESSING            = C.D3DCREATE_HARDWARE_VERTEXPROCESSING
	CREATE_MIXED_VERTEXPROCESSING               = C.D3DCREATE_MIXED_VERTEXPROCESSING
	CREATE_MULTITHREADED                        = C.D3DCREATE_MULTITHREADED
	CREATE_NOWINDOWCHANGES                      = 0x00000800
	CREATE_PUREDEVICE                           = C.D3DCREATE_PUREDEVICE
	CREATE_SOFTWARE_VERTEXPROCESSING            = C.D3DCREATE_SOFTWARE_VERTEXPROCESSING
	CS_ALL                                      = C.D3DCS_ALL
	CS_LEFT                                     = C.D3DCS_LEFT
	CS_RIGHT                                    = C.D3DCS_RIGHT
	CS_TOP                                      = C.D3DCS_TOP
	CS_BOTTOM                                   = C.D3DCS_BOTTOM
	CS_FRONT                                    = C.D3DCS_FRONT
	CS_BACK                                     = C.D3DCS_BACK
	CS_PLANE0                                   = C.D3DCS_PLANE0
	CS_PLANE1                                   = C.D3DCS_PLANE1
	CS_PLANE2                                   = C.D3DCS_PLANE2
	CS_PLANE3                                   = C.D3DCS_PLANE3
	CS_PLANE4                                   = C.D3DCS_PLANE4
	CS_PLANE5                                   = C.D3DCS_PLANE5
	CURSORCAPS_COLOR                            = C.D3DCURSORCAPS_COLOR
	CURSORCAPS_LOWRES                           = C.D3DCURSORCAPS_LOWRES
	DEVCAPS_CANBLTSYSTONONLOCAL                 = C.D3DDEVCAPS_CANBLTSYSTONONLOCAL
	DEVCAPS_CANRENDERAFTERFLIP                  = C.D3DDEVCAPS_CANRENDERAFTERFLIP
	DEVCAPS_DRAWPRIMITIVES2                     = C.D3DDEVCAPS_DRAWPRIMITIVES2
	DEVCAPS_DRAWPRIMITIVES2EX                   = C.D3DDEVCAPS_DRAWPRIMITIVES2EX
	DEVCAPS_DRAWPRIMTLVERTEX                    = C.D3DDEVCAPS_DRAWPRIMTLVERTEX
	DEVCAPS_EXECUTESYSTEMMEMORY                 = C.D3DDEVCAPS_EXECUTESYSTEMMEMORY
	DEVCAPS_EXECUTEVIDEOMEMORY                  = C.D3DDEVCAPS_EXECUTEVIDEOMEMORY
	DEVCAPS_HWRASTERIZATION                     = C.D3DDEVCAPS_HWRASTERIZATION
	DEVCAPS_HWTRANSFORMANDLIGHT                 = C.D3DDEVCAPS_HWTRANSFORMANDLIGHT
	DEVCAPS_NPATCHES                            = C.D3DDEVCAPS_NPATCHES
	DEVCAPS_PUREDEVICE                          = C.D3DDEVCAPS_PUREDEVICE
	DEVCAPS_QUINTICRTPATCHES                    = C.D3DDEVCAPS_QUINTICRTPATCHES
	DEVCAPS_RTPATCHES                           = C.D3DDEVCAPS_RTPATCHES
	DEVCAPS_RTPATCHHANDLEZERO                   = C.D3DDEVCAPS_RTPATCHHANDLEZERO
	DEVCAPS_SEPARATETEXTUREMEMORIES             = C.D3DDEVCAPS_SEPARATETEXTUREMEMORIES
	DEVCAPS_TEXTURENONLOCALVIDMEM               = C.D3DDEVCAPS_TEXTURENONLOCALVIDMEM
	DEVCAPS_TEXTURESYSTEMMEMORY                 = C.D3DDEVCAPS_TEXTURESYSTEMMEMORY
	DEVCAPS_TEXTUREVIDEOMEMORY                  = C.D3DDEVCAPS_TEXTUREVIDEOMEMORY
	DEVCAPS_TLVERTEXSYSTEMMEMORY                = C.D3DDEVCAPS_TLVERTEXSYSTEMMEMORY
	DEVCAPS_TLVERTEXVIDEOMEMORY                 = C.D3DDEVCAPS_TLVERTEXVIDEOMEMORY
	DEVCAPS2_ADAPTIVETESSRTPATCH                = C.D3DDEVCAPS2_ADAPTIVETESSRTPATCH
	DEVCAPS2_ADAPTIVETESSNPATCH                 = C.D3DDEVCAPS2_ADAPTIVETESSNPATCH
	DEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES      = C.D3DDEVCAPS2_CAN_STRETCHRECT_FROM_TEXTURES
	DEVCAPS2_DMAPNPATCH                         = C.D3DDEVCAPS2_DMAPNPATCH
	DEVCAPS2_PRESAMPLEDDMAPNPATCH               = C.D3DDEVCAPS2_PRESAMPLEDDMAPNPATCH
	DEVCAPS2_STREAMOFFSET                       = C.D3DDEVCAPS2_STREAMOFFSET
	DEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET = C.D3DDEVCAPS2_VERTEXELEMENTSCANSHARESTREAMOFFSET
	DTCAPS_UBYTE4                               = C.D3DDTCAPS_UBYTE4
	DTCAPS_UBYTE4N                              = C.D3DDTCAPS_UBYTE4N
	DTCAPS_SHORT2N                              = C.D3DDTCAPS_SHORT2N
	DTCAPS_SHORT4N                              = C.D3DDTCAPS_SHORT4N
	DTCAPS_USHORT2N                             = C.D3DDTCAPS_USHORT2N
	DTCAPS_USHORT4N                             = C.D3DDTCAPS_USHORT4N
	DTCAPS_UDEC3                                = C.D3DDTCAPS_UDEC3
	DTCAPS_DEC3N                                = C.D3DDTCAPS_DEC3N
	DTCAPS_FLOAT16_2                            = C.D3DDTCAPS_FLOAT16_2
	DTCAPS_FLOAT16_4                            = C.D3DDTCAPS_FLOAT16_4
	OK_NOAUTOGEN                                = C.D3DOK_NOAUTOGEN
	ERR_CONFLICTINGRENDERSTATE                  = C.D3DERR_CONFLICTINGRENDERSTATE
	ERR_CONFLICTINGTEXTUREFILTER                = C.D3DERR_CONFLICTINGTEXTUREFILTER
	ERR_CONFLICTINGTEXTUREPALETTE               = C.D3DERR_CONFLICTINGTEXTUREPALETTE
	ERR_DEVICEHUNG                              = -2005530508
	ERR_DEVICELOST                              = C.D3DERR_DEVICELOST
	ERR_DEVICENOTRESET                          = C.D3DERR_DEVICENOTRESET
	ERR_DEVICEREMOVED                           = -2005530512
	ERR_DRIVERINTERNALERROR                     = C.D3DERR_DRIVERINTERNALERROR
	ERR_DRIVERINVALIDCALL                       = C.D3DERR_DRIVERINVALIDCALL
	ERR_INVALIDCALL                             = C.D3DERR_INVALIDCALL
	ERR_INVALIDDEVICE                           = C.D3DERR_INVALIDDEVICE
	ERR_MOREDATA                                = C.D3DERR_MOREDATA
	ERR_NOTAVAILABLE                            = C.D3DERR_NOTAVAILABLE
	ERR_NOTFOUND                                = C.D3DERR_NOTFOUND
	OK                                          = C.D3D_OK
	ERR_OUTOFVIDEOMEMORY                        = C.D3DERR_OUTOFVIDEOMEMORY
	ERR_TOOMANYOPERATIONS                       = C.D3DERR_TOOMANYOPERATIONS
	ERR_UNSUPPORTEDALPHAARG                     = C.D3DERR_UNSUPPORTEDALPHAARG
	ERR_UNSUPPORTEDALPHAOPERATION               = C.D3DERR_UNSUPPORTEDALPHAOPERATION
	ERR_UNSUPPORTEDCOLORARG                     = C.D3DERR_UNSUPPORTEDCOLORARG
	ERR_UNSUPPORTEDCOLOROPERATION               = C.D3DERR_UNSUPPORTEDCOLOROPERATION
	ERR_UNSUPPORTEDFACTORVALUE                  = C.D3DERR_UNSUPPORTEDFACTORVALUE
	ERR_UNSUPPORTEDTEXTUREFILTER                = C.D3DERR_UNSUPPORTEDTEXTUREFILTER
	ERR_WASSTILLDRAWING                         = C.D3DERR_WASSTILLDRAWING
	ERR_WRONGTEXTUREFORMAT                      = C.D3DERR_WRONGTEXTUREFORMAT
	E_FAIL                                      = C.E_FAIL
	E_INVALIDARG                                = C.E_INVALIDARG
	E_NOINTERFACE                               = C.E_NOINTERFACE
	E_NOTIMPL                                   = C.E_NOTIMPL
	E_OUTOFMEMORY                               = C.E_OUTOFMEMORY
	S_NOT_RESIDENT                              = 141953141
	S_RESIDENT_IN_SHARED_MEMORY                 = 141953142
	ERR_UNSUPPORTEDOVERLAY                      = -2005530501
	ERR_UNSUPPORTEDOVERLAYFORMAT                = -2005530500
	ERR_CANNOTPROTECTCONTENT                    = -2005530499
	ERR_UNSUPPORTEDCRYPTO                       = -2005530498
	ERR_PRESENT_STATISTICS_DISJOINT             = -2005530492
	FVFCAPS_DONOTSTRIPELEMENTS                  = C.D3DFVFCAPS_DONOTSTRIPELEMENTS
	FVFCAPS_PSIZE                               = C.D3DFVFCAPS_PSIZE
	FVFCAPS_TEXCOORDCOUNTMASK                   = C.D3DFVFCAPS_TEXCOORDCOUNTMASK
	FVF_DIFFUSE                                 = C.D3DFVF_DIFFUSE
	FVF_NORMAL                                  = C.D3DFVF_NORMAL
	FVF_PSIZE                                   = C.D3DFVF_PSIZE
	FVF_SPECULAR                                = C.D3DFVF_SPECULAR
	FVF_XYZ                                     = C.D3DFVF_XYZ
	FVF_XYZRHW                                  = C.D3DFVF_XYZRHW
	FVF_XYZB1                                   = C.D3DFVF_XYZB1
	FVF_XYZB2                                   = C.D3DFVF_XYZB2
	FVF_XYZB3                                   = C.D3DFVF_XYZB3
	FVF_XYZB4                                   = C.D3DFVF_XYZB4
	FVF_XYZB5                                   = C.D3DFVF_XYZB5
	FVF_XYZW                                    = C.D3DFVF_XYZW
	FVF_TEX0                                    = C.D3DFVF_TEX0
	FVF_TEX1                                    = C.D3DFVF_TEX1
	FVF_TEX2                                    = C.D3DFVF_TEX2
	FVF_TEX3                                    = C.D3DFVF_TEX3
	FVF_TEX4                                    = C.D3DFVF_TEX4
	FVF_TEX5                                    = C.D3DFVF_TEX5
	FVF_TEX6                                    = C.D3DFVF_TEX6
	FVF_TEX7                                    = C.D3DFVF_TEX7
	FVF_TEX8                                    = C.D3DFVF_TEX8
	FVF_TEXTUREFORMAT1                          = C.D3DFVF_TEXTUREFORMAT1
	FVF_TEXTUREFORMAT2                          = C.D3DFVF_TEXTUREFORMAT2
	FVF_TEXTUREFORMAT3                          = C.D3DFVF_TEXTUREFORMAT3
	FVF_TEXTUREFORMAT4                          = C.D3DFVF_TEXTUREFORMAT4
	FVF_POSITION_MASK                           = C.D3DFVF_POSITION_MASK
	FVF_TEXCOUNT_MASK                           = C.D3DFVF_TEXCOUNT_MASK
	FVF_LASTBETA_D3DCOLOR                       = C.D3DFVF_LASTBETA_D3DCOLOR
	FVF_LASTBETA_UBYTE4                         = C.D3DFVF_LASTBETA_UBYTE4
	FVF_TEXCOUNT_SHIFT                          = C.D3DFVF_TEXCOUNT_SHIFT
	LINECAPS_ALPHACMP                           = C.D3DLINECAPS_ALPHACMP
	LINECAPS_ANTIALIAS                          = C.D3DLINECAPS_ANTIALIAS
	LINECAPS_BLEND                              = C.D3DLINECAPS_BLEND
	LINECAPS_FOG                                = C.D3DLINECAPS_FOG
	LINECAPS_TEXTURE                            = C.D3DLINECAPS_TEXTURE
	LINECAPS_ZTEST                              = C.D3DLINECAPS_ZTEST
	LOCK_DISCARD                                = C.D3DLOCK_DISCARD
	LOCK_DONOTWAIT                              = C.D3DLOCK_DONOTWAIT
	LOCK_NO_DIRTY_UPDATE                        = C.D3DLOCK_NO_DIRTY_UPDATE
	LOCK_NOOVERWRITE                            = C.D3DLOCK_NOOVERWRITE
	LOCK_NOSYSLOCK                              = C.D3DLOCK_NOSYSLOCK
	LOCK_READONLY                               = C.D3DLOCK_READONLY
	PBLENDCAPS_BLENDFACTOR                      = C.D3DPBLENDCAPS_BLENDFACTOR
	PBLENDCAPS_BOTHINVSRCALPHA                  = C.D3DPBLENDCAPS_BOTHINVSRCALPHA
	PBLENDCAPS_BOTHSRCALPHA                     = C.D3DPBLENDCAPS_BOTHSRCALPHA
	PBLENDCAPS_DESTALPHA                        = C.D3DPBLENDCAPS_DESTALPHA
	PBLENDCAPS_DESTCOLOR                        = C.D3DPBLENDCAPS_DESTCOLOR
	PBLENDCAPS_INVDESTALPHA                     = C.D3DPBLENDCAPS_INVDESTALPHA
	PBLENDCAPS_INVDESTCOLOR                     = C.D3DPBLENDCAPS_INVDESTCOLOR
	PBLENDCAPS_INVSRCALPHA                      = C.D3DPBLENDCAPS_INVSRCALPHA
	PBLENDCAPS_INVSRCCOLOR                      = C.D3DPBLENDCAPS_INVSRCCOLOR
	PBLENDCAPS_ONE                              = C.D3DPBLENDCAPS_ONE
	PBLENDCAPS_SRCALPHA                         = C.D3DPBLENDCAPS_SRCALPHA
	PBLENDCAPS_SRCALPHASAT                      = C.D3DPBLENDCAPS_SRCALPHASAT
	PBLENDCAPS_SRCCOLOR                         = C.D3DPBLENDCAPS_SRCCOLOR
	PBLENDCAPS_ZERO                             = C.D3DPBLENDCAPS_ZERO
	PCMPCAPS_ALWAYS                             = C.D3DPCMPCAPS_ALWAYS
	PCMPCAPS_EQUAL                              = C.D3DPCMPCAPS_EQUAL
	PCMPCAPS_GREATER                            = C.D3DPCMPCAPS_GREATER
	PCMPCAPS_GREATEREQUAL                       = C.D3DPCMPCAPS_GREATEREQUAL
	PCMPCAPS_LESS                               = C.D3DPCMPCAPS_LESS
	PCMPCAPS_LESSEQUAL                          = C.D3DPCMPCAPS_LESSEQUAL
	PCMPCAPS_NEVER                              = C.D3DPCMPCAPS_NEVER
	PCMPCAPS_NOTEQUAL                           = C.D3DPCMPCAPS_NOTEQUAL
	PMISCCAPS_MASKZ                             = C.D3DPMISCCAPS_MASKZ
	PMISCCAPS_CULLNONE                          = C.D3DPMISCCAPS_CULLNONE
	PMISCCAPS_CULLCW                            = C.D3DPMISCCAPS_CULLCW
	PMISCCAPS_CULLCCW                           = C.D3DPMISCCAPS_CULLCCW
	PMISCCAPS_COLORWRITEENABLE                  = C.D3DPMISCCAPS_COLORWRITEENABLE
	PMISCCAPS_CLIPPLANESCALEDPOINTS             = C.D3DPMISCCAPS_CLIPPLANESCALEDPOINTS
	PMISCCAPS_CLIPTLVERTS                       = C.D3DPMISCCAPS_CLIPTLVERTS
	PMISCCAPS_TSSARGTEMP                        = C.D3DPMISCCAPS_TSSARGTEMP
	PMISCCAPS_BLENDOP                           = C.D3DPMISCCAPS_BLENDOP
	PMISCCAPS_NULLREFERENCE                     = C.D3DPMISCCAPS_NULLREFERENCE
	PMISCCAPS_INDEPENDENTWRITEMASKS             = C.D3DPMISCCAPS_INDEPENDENTWRITEMASKS
	PMISCCAPS_PERSTAGECONSTANT                  = C.D3DPMISCCAPS_PERSTAGECONSTANT
	PMISCCAPS_FOGANDSPECULARALPHA               = C.D3DPMISCCAPS_FOGANDSPECULARALPHA
	PMISCCAPS_SEPARATEALPHABLEND                = C.D3DPMISCCAPS_SEPARATEALPHABLEND
	PMISCCAPS_MRTINDEPENDENTBITDEPTHS           = C.D3DPMISCCAPS_MRTINDEPENDENTBITDEPTHS
	PMISCCAPS_MRTPOSTPIXELSHADERBLENDING        = C.D3DPMISCCAPS_MRTPOSTPIXELSHADERBLENDING
	PMISCCAPS_FOGVERTEXCLAMPED                  = C.D3DPMISCCAPS_FOGVERTEXCLAMPED
	PRASTERCAPS_ANISOTROPY                      = C.D3DPRASTERCAPS_ANISOTROPY
	PRASTERCAPS_COLORPERSPECTIVE                = C.D3DPRASTERCAPS_COLORPERSPECTIVE
	PRASTERCAPS_DITHER                          = C.D3DPRASTERCAPS_DITHER
	PRASTERCAPS_DEPTHBIAS                       = C.D3DPRASTERCAPS_DEPTHBIAS
	PRASTERCAPS_FOGRANGE                        = C.D3DPRASTERCAPS_FOGRANGE
	PRASTERCAPS_FOGTABLE                        = C.D3DPRASTERCAPS_FOGTABLE
	PRASTERCAPS_FOGVERTEX                       = C.D3DPRASTERCAPS_FOGVERTEX
	PRASTERCAPS_MIPMAPLODBIAS                   = C.D3DPRASTERCAPS_MIPMAPLODBIAS
	PRASTERCAPS_MULTISAMPLE_TOGGLE              = C.D3DPRASTERCAPS_MULTISAMPLE_TOGGLE
	PRASTERCAPS_SCISSORTEST                     = C.D3DPRASTERCAPS_SCISSORTEST
	PRASTERCAPS_SLOPESCALEDEPTHBIAS             = C.D3DPRASTERCAPS_SLOPESCALEDEPTHBIAS
	PRASTERCAPS_WBUFFER                         = C.D3DPRASTERCAPS_WBUFFER
	PRASTERCAPS_WFOG                            = C.D3DPRASTERCAPS_WFOG
	PRASTERCAPS_ZBUFFERLESSHSR                  = C.D3DPRASTERCAPS_ZBUFFERLESSHSR
	PRASTERCAPS_ZFOG                            = C.D3DPRASTERCAPS_ZFOG
	PRASTERCAPS_ZTEST                           = C.D3DPRASTERCAPS_ZTEST
	PRESENT_INTERVAL_DEFAULT                    = C.D3DPRESENT_INTERVAL_DEFAULT
	PRESENT_INTERVAL_ONE                        = C.D3DPRESENT_INTERVAL_ONE
	PRESENT_INTERVAL_TWO                        = C.D3DPRESENT_INTERVAL_TWO
	PRESENT_INTERVAL_THREE                      = C.D3DPRESENT_INTERVAL_THREE
	PRESENT_INTERVAL_FOUR                       = C.D3DPRESENT_INTERVAL_FOUR
	PRESENT_INTERVAL_IMMEDIATE                  = C.D3DPRESENT_INTERVAL_IMMEDIATE
	PRESENT_BACK_BUFFERS_MAX                    = C.D3DPRESENT_BACK_BUFFERS_MAX
	PRESENTFLAG_DEVICECLIP                      = C.D3DPRESENTFLAG_DEVICECLIP
	PRESENTFLAG_DISCARD_DEPTHSTENCIL            = C.D3DPRESENTFLAG_DISCARD_DEPTHSTENCIL
	PRESENTFLAG_LOCKABLE_BACKBUFFER             = C.D3DPRESENTFLAG_LOCKABLE_BACKBUFFER
	PRESENTFLAG_NOAUTOROTATE                    = 32
	PRESENTFLAG_UNPRUNEDMODE                    = 64
	PRESENTFLAG_VIDEO                           = C.D3DPRESENTFLAG_VIDEO
	PSHADECAPS_ALPHAGOURAUDBLEND                = C.D3DPSHADECAPS_ALPHAGOURAUDBLEND
	PSHADECAPS_COLORGOURAUDRGB                  = C.D3DPSHADECAPS_COLORGOURAUDRGB
	PSHADECAPS_FOGGOURAUD                       = C.D3DPSHADECAPS_FOGGOURAUD
	PSHADECAPS_SPECULARGOURAUDRGB               = C.D3DPSHADECAPS_SPECULARGOURAUDRGB
	PS20_MAX_DYNAMICFLOWCONTROLDEPTH            = C.D3DPS20_MAX_DYNAMICFLOWCONTROLDEPTH
	PS20_MIN_DYNAMICFLOWCONTROLDEPTH            = C.D3DPS20_MIN_DYNAMICFLOWCONTROLDEPTH
	PS20_MAX_NUMTEMPS                           = C.D3DPS20_MAX_NUMTEMPS
	PS20_MIN_NUMTEMPS                           = C.D3DPS20_MIN_NUMTEMPS
	PS20_MAX_STATICFLOWCONTROLDEPTH             = C.D3DPS20_MAX_STATICFLOWCONTROLDEPTH
	PS20_MIN_STATICFLOWCONTROLDEPTH             = C.D3DPS20_MIN_STATICFLOWCONTROLDEPTH
	PS20_MAX_NUMINSTRUCTIONSLOTS                = C.D3DPS20_MAX_NUMINSTRUCTIONSLOTS
	PS20_MIN_NUMINSTRUCTIONSLOTS                = C.D3DPS20_MIN_NUMINSTRUCTIONSLOTS
	PTADDRESSCAPS_BORDER                        = C.D3DPTADDRESSCAPS_BORDER
	PTADDRESSCAPS_CLAMP                         = C.D3DPTADDRESSCAPS_CLAMP
	PTADDRESSCAPS_INDEPENDENTUV                 = C.D3DPTADDRESSCAPS_INDEPENDENTUV
	PTADDRESSCAPS_MIRROR                        = C.D3DPTADDRESSCAPS_MIRROR
	PTADDRESSCAPS_MIRRORONCE                    = C.D3DPTADDRESSCAPS_MIRRORONCE
	PTADDRESSCAPS_WRAP                          = C.D3DPTADDRESSCAPS_WRAP
	PTEXTURECAPS_ALPHA                          = C.D3DPTEXTURECAPS_ALPHA
	PTEXTURECAPS_ALPHAPALETTE                   = C.D3DPTEXTURECAPS_ALPHAPALETTE
	PTEXTURECAPS_CUBEMAP                        = C.D3DPTEXTURECAPS_CUBEMAP
	PTEXTURECAPS_CUBEMAP_POW2                   = C.D3DPTEXTURECAPS_CUBEMAP_POW2
	PTEXTURECAPS_MIPCUBEMAP                     = C.D3DPTEXTURECAPS_MIPCUBEMAP
	PTEXTURECAPS_MIPMAP                         = C.D3DPTEXTURECAPS_MIPMAP
	PTEXTURECAPS_MIPVOLUMEMAP                   = C.D3DPTEXTURECAPS_MIPVOLUMEMAP
	PTEXTURECAPS_NONPOW2CONDITIONAL             = C.D3DPTEXTURECAPS_NONPOW2CONDITIONAL
	PTEXTURECAPS_NOPROJECTEDBUMPENV             = C.D3DPTEXTURECAPS_NOPROJECTEDBUMPENV
	PTEXTURECAPS_PERSPECTIVE                    = C.D3DPTEXTURECAPS_PERSPECTIVE
	PTEXTURECAPS_POW2                           = C.D3DPTEXTURECAPS_POW2
	PTEXTURECAPS_PROJECTED                      = C.D3DPTEXTURECAPS_PROJECTED
	PTEXTURECAPS_SQUAREONLY                     = C.D3DPTEXTURECAPS_SQUAREONLY
	PTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE       = C.D3DPTEXTURECAPS_TEXREPEATNOTSCALEDBYSIZE
	PTEXTURECAPS_VOLUMEMAP                      = C.D3DPTEXTURECAPS_VOLUMEMAP
	PTEXTURECAPS_VOLUMEMAP_POW2                 = C.D3DPTEXTURECAPS_VOLUMEMAP_POW2
	PTFILTERCAPS_MAGFPOINT                      = C.D3DPTFILTERCAPS_MAGFPOINT
	PTFILTERCAPS_MAGFLINEAR                     = C.D3DPTFILTERCAPS_MAGFLINEAR
	PTFILTERCAPS_MAGFANISOTROPIC                = C.D3DPTFILTERCAPS_MAGFANISOTROPIC
	PTFILTERCAPS_MAGFPYRAMIDALQUAD              = C.D3DPTFILTERCAPS_MAGFPYRAMIDALQUAD
	PTFILTERCAPS_MAGFGAUSSIANQUAD               = C.D3DPTFILTERCAPS_MAGFGAUSSIANQUAD
	PTFILTERCAPS_MINFPOINT                      = C.D3DPTFILTERCAPS_MINFPOINT
	PTFILTERCAPS_MINFLINEAR                     = C.D3DPTFILTERCAPS_MINFLINEAR
	PTFILTERCAPS_MINFANISOTROPIC                = C.D3DPTFILTERCAPS_MINFANISOTROPIC
	PTFILTERCAPS_MINFPYRAMIDALQUAD              = C.D3DPTFILTERCAPS_MINFPYRAMIDALQUAD
	PTFILTERCAPS_MINFGAUSSIANQUAD               = C.D3DPTFILTERCAPS_MINFGAUSSIANQUAD
	PTFILTERCAPS_MIPFPOINT                      = C.D3DPTFILTERCAPS_MIPFPOINT
	PTFILTERCAPS_MIPFLINEAR                     = C.D3DPTFILTERCAPS_MIPFLINEAR
	SPD_IUNKNOWN                                = C.D3DSPD_IUNKNOWN
	STENCILCAPS_KEEP                            = C.D3DSTENCILCAPS_KEEP
	STENCILCAPS_ZERO                            = C.D3DSTENCILCAPS_ZERO
	STENCILCAPS_REPLACE                         = C.D3DSTENCILCAPS_REPLACE
	STENCILCAPS_INCRSAT                         = C.D3DSTENCILCAPS_INCRSAT
	STENCILCAPS_DECRSAT                         = C.D3DSTENCILCAPS_DECRSAT
	STENCILCAPS_INVERT                          = C.D3DSTENCILCAPS_INVERT
	STENCILCAPS_INCR                            = C.D3DSTENCILCAPS_INCR
	STENCILCAPS_DECR                            = C.D3DSTENCILCAPS_DECR
	STENCILCAPS_TWOSIDED                        = C.D3DSTENCILCAPS_TWOSIDED
	TA_CONSTANT                                 = C.D3DTA_CONSTANT
	TA_CURRENT                                  = C.D3DTA_CURRENT
	TA_DIFFUSE                                  = C.D3DTA_DIFFUSE
	TA_SELECTMASK                               = C.D3DTA_SELECTMASK
	TA_SPECULAR                                 = C.D3DTA_SPECULAR
	TA_TEMP                                     = C.D3DTA_TEMP
	TA_TEXTURE                                  = C.D3DTA_TEXTURE
	TA_TFACTOR                                  = C.D3DTA_TFACTOR
	TA_ALPHAREPLICATE                           = C.D3DTA_ALPHAREPLICATE
	TA_COMPLEMENT                               = C.D3DTA_COMPLEMENT
	TEXOPCAPS_ADD                               = C.D3DTEXOPCAPS_ADD
	TEXOPCAPS_ADDSIGNED                         = C.D3DTEXOPCAPS_ADDSIGNED
	TEXOPCAPS_ADDSIGNED2X                       = C.D3DTEXOPCAPS_ADDSIGNED2X
	TEXOPCAPS_ADDSMOOTH                         = C.D3DTEXOPCAPS_ADDSMOOTH
	TEXOPCAPS_BLENDCURRENTALPHA                 = C.D3DTEXOPCAPS_BLENDCURRENTALPHA
	TEXOPCAPS_BLENDDIFFUSEALPHA                 = C.D3DTEXOPCAPS_BLENDDIFFUSEALPHA
	TEXOPCAPS_BLENDFACTORALPHA                  = C.D3DTEXOPCAPS_BLENDFACTORALPHA
	TEXOPCAPS_BLENDTEXTUREALPHA                 = C.D3DTEXOPCAPS_BLENDTEXTUREALPHA
	TEXOPCAPS_BLENDTEXTUREALPHAPM               = C.D3DTEXOPCAPS_BLENDTEXTUREALPHAPM
	TEXOPCAPS_BUMPENVMAP                        = C.D3DTEXOPCAPS_BUMPENVMAP
	TEXOPCAPS_BUMPENVMAPLUMINANCE               = C.D3DTEXOPCAPS_BUMPENVMAPLUMINANCE
	TEXOPCAPS_DISABLE                           = C.D3DTEXOPCAPS_DISABLE
	TEXOPCAPS_DOTPRODUCT3                       = C.D3DTEXOPCAPS_DOTPRODUCT3
	TEXOPCAPS_LERP                              = C.D3DTEXOPCAPS_LERP
	TEXOPCAPS_MODULATE                          = C.D3DTEXOPCAPS_MODULATE
	TEXOPCAPS_MODULATE2X                        = C.D3DTEXOPCAPS_MODULATE2X
	TEXOPCAPS_MODULATE4X                        = C.D3DTEXOPCAPS_MODULATE4X
	TEXOPCAPS_MODULATEALPHA_ADDCOLOR            = C.D3DTEXOPCAPS_MODULATEALPHA_ADDCOLOR
	TEXOPCAPS_MODULATECOLOR_ADDALPHA            = C.D3DTEXOPCAPS_MODULATECOLOR_ADDALPHA
	TEXOPCAPS_MODULATEINVALPHA_ADDCOLOR         = C.D3DTEXOPCAPS_MODULATEINVALPHA_ADDCOLOR
	TEXOPCAPS_MODULATEINVCOLOR_ADDALPHA         = C.D3DTEXOPCAPS_MODULATEINVCOLOR_ADDALPHA
	TEXOPCAPS_MULTIPLYADD                       = C.D3DTEXOPCAPS_MULTIPLYADD
	TEXOPCAPS_PREMODULATE                       = C.D3DTEXOPCAPS_PREMODULATE
	TEXOPCAPS_SELECTARG1                        = C.D3DTEXOPCAPS_SELECTARG1
	TEXOPCAPS_SELECTARG2                        = C.D3DTEXOPCAPS_SELECTARG2
	TEXOPCAPS_SUBTRACT                          = C.D3DTEXOPCAPS_SUBTRACT
	TSS_TCI_PASSTHRU                            = C.D3DTSS_TCI_PASSTHRU
	TSS_TCI_CAMERASPACENORMAL                   = C.D3DTSS_TCI_CAMERASPACENORMAL
	TSS_TCI_CAMERASPACEPOSITION                 = C.D3DTSS_TCI_CAMERASPACEPOSITION
	TSS_TCI_CAMERASPACEREFLECTIONVECTOR         = C.D3DTSS_TCI_CAMERASPACEREFLECTIONVECTOR
	TSS_TCI_SPHEREMAP                           = C.D3DTSS_TCI_SPHEREMAP
	USAGE_AUTOGENMIPMAP                         = C.D3DUSAGE_AUTOGENMIPMAP
	USAGE_DEPTHSTENCIL                          = C.D3DUSAGE_DEPTHSTENCIL
	USAGE_DMAP                                  = C.D3DUSAGE_DMAP
	USAGE_DONOTCLIP                             = C.D3DUSAGE_DONOTCLIP
	USAGE_DYNAMIC                               = C.D3DUSAGE_DYNAMIC
	USAGE_NPATCHES                              = C.D3DUSAGE_NPATCHES
	USAGE_POINTS                                = C.D3DUSAGE_POINTS
	USAGE_RENDERTARGET                          = C.D3DUSAGE_RENDERTARGET
	USAGE_RTPATCHES                             = C.D3DUSAGE_RTPATCHES
	USAGE_SOFTWAREPROCESSING                    = C.D3DUSAGE_SOFTWAREPROCESSING
	USAGE_WRITEONLY                             = C.D3DUSAGE_WRITEONLY
	USAGE_QUERY_FILTER                          = C.D3DUSAGE_QUERY_FILTER
	USAGE_QUERY_LEGACYBUMPMAP                   = C.D3DUSAGE_QUERY_LEGACYBUMPMAP
	USAGE_QUERY_POSTPIXELSHADER_BLENDING        = C.D3DUSAGE_QUERY_POSTPIXELSHADER_BLENDING
	USAGE_QUERY_SRGBREAD                        = C.D3DUSAGE_QUERY_SRGBREAD
	USAGE_QUERY_SRGBWRITE                       = C.D3DUSAGE_QUERY_SRGBWRITE
	USAGE_QUERY_VERTEXTEXTURE                   = C.D3DUSAGE_QUERY_VERTEXTEXTURE
	USAGE_QUERY_WRAPANDMIP                      = 2097152
	VERTEXTEXTURESAMPLER0                       = C.D3DVERTEXTEXTURESAMPLER0
	VERTEXTEXTURESAMPLER1                       = C.D3DVERTEXTEXTURESAMPLER1
	VERTEXTEXTURESAMPLER2                       = C.D3DVERTEXTEXTURESAMPLER2
	VERTEXTEXTURESAMPLER3                       = C.D3DVERTEXTEXTURESAMPLER3
	VS20CAPS_PREDICATION                        = C.D3DVS20CAPS_PREDICATION
	VS20_MAX_DYNAMICFLOWCONTROLDEPTH            = C.D3DVS20_MAX_DYNAMICFLOWCONTROLDEPTH
	VS20_MIN_DYNAMICFLOWCONTROLDEPTH            = C.D3DVS20_MIN_DYNAMICFLOWCONTROLDEPTH
	VS20_MAX_NUMTEMPS                           = C.D3DVS20_MAX_NUMTEMPS
	VS20_MIN_NUMTEMPS                           = C.D3DVS20_MIN_NUMTEMPS
	VS20_MAX_STATICFLOWCONTROLDEPTH             = C.D3DVS20_MAX_STATICFLOWCONTROLDEPTH
	VS20_MIN_STATICFLOWCONTROLDEPTH             = C.D3DVS20_MIN_STATICFLOWCONTROLDEPTH
	VTXPCAPS_DIRECTIONALLIGHTS                  = C.D3DVTXPCAPS_DIRECTIONALLIGHTS
	VTXPCAPS_LOCALVIEWER                        = C.D3DVTXPCAPS_LOCALVIEWER
	VTXPCAPS_MATERIALSOURCE7                    = C.D3DVTXPCAPS_MATERIALSOURCE7
	VTXPCAPS_NO_TEXGEN_NONLOCALVIEWER           = C.D3DVTXPCAPS_NO_TEXGEN_NONLOCALVIEWER
	VTXPCAPS_POSITIONALLIGHTS                   = C.D3DVTXPCAPS_POSITIONALLIGHTS
	VTXPCAPS_TEXGEN                             = C.D3DVTXPCAPS_TEXGEN
	VTXPCAPS_TEXGEN_SPHEREMAP                   = C.D3DVTXPCAPS_TEXGEN_SPHEREMAP
	VTXPCAPS_TWEENING                           = C.D3DVTXPCAPS_TWEENING
	S_OK                                        = C.S_OK
	S_PRESENT_OCCLUDED                          = 141953144
	S_PRESENT_MODE_CHANGED                      = 141953143
	SDK_VERSION                                 = C.D3D_SDK_VERSION
	MAX_SIMULTANEOUS_RENDERTARGETS              = C.D3D_MAX_SIMULTANEOUS_RENDERTARGETS
	DMAPSAMPLER                                 = C.D3DDMAPSAMPLER
	DP_MAXTEXCOORD                              = C.D3DDP_MAXTEXCOORD
	MAXD3DDECLLENGTH                            = C.MAXD3DDECLLENGTH
	MAXD3DDECLUSAGEINDEX                        = C.MAXD3DDECLUSAGEINDEX
)
