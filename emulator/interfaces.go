package emu

import (
	_ "github.com/aniou/morfe/emulator/vram"
)

var (
	DIP = [9]bool{false}		   // there are 8 switches, but numbering from 1 is convinient
	                                   // thus DIP[0] is never used
)

type Processor interface {
        Reset()
        Execute() uint32	           // execute one or more steps and returns used cycles
        Step() uint32		           // execute single step and returns used cycles
	GetRegisters() map[string]uint32   // returns all registers of CPU
	GetType() uint			   // returns CPU id (when many types are available)
	IsEnabled() bool                   // as name suggests
	Enable(bool)			   // enables/disables CPU
	Dissasm() string
	GetCycles() uint32		   // number of cycles used by last step
	GetAllCycles() uint64	           // cumulative number of cycles used
	StatusString() string		   // string that represents status flags
	ResetCycles()
	TriggerIRQ()
	SetRegister(string, uint32) error  // set selected register
	SetPC(uint32)			   // redundant to SetRegister but convinient

	Write_8(uint32, byte)		   // write byte to   cpu memory
	Read_8(uint32) byte                // read  byte from cpu memory

	GetName() string	           // get id as "cpu0" / "cpu1" of unit
	DisassembleCurrentPC() string	   // disassemble current line
}

type Bus interface {
	Write_8(byte, uint32, byte)
	Read_8 (byte, uint32) byte
}

type Memory interface {
        Write(fn byte, addr uint32, value byte)  error
        Read (fn byte, addr uint32)             (byte, error)
	Name(fn byte) string
        Size(fn byte) (uint32, uint32)

        //Shutdown()
        //Clear()
        //Dump(address uint32) []byte
}

type GPU interface {
        Write(fn byte, addr uint32, value byte)  error
        Read (fn byte, addr uint32)             (byte, error)
	Name(fn byte) string
        Size(fn byte) (uint32, uint32)

	GetCommon() *GPU_common
	RenderBitmapText()
}

const (
        CPU_65c816 = 0
        CPU_68000  = 1
        CPU_68030  = 2
)

const (
	M_USER  = 0
	M_SV    = 1
)

// a 'common' set of Vicky's data
type GPU_common struct {
        TFB     []uint32       // text   framebuffer
        BM0FB   []uint32       // bitmap0 framebuffer
        BM1FB   []uint32       // bitmap1 framebuffer

        // some convinient registers that should be converted
        // into some kind of memory indexes...
        Master_L        byte    // MASTER_CTRL_REG_L
        Master_H        byte    // MASTER_CTRL_REG_H
        Cursor_visible  bool
        Border_visible  bool
        BM0_visible     bool
        BM1_visible     bool

        Border_color_b  byte
        Border_color_g  byte
        Border_color_r  byte
        Border_x_size   int32
        Border_y_size   int32
        Background      [3]byte         // r, g, b
}

