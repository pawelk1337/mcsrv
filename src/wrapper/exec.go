package wrapper

import (
	"fmt"
	"io"
	"os/exec"
)

type JavaExec interface {
	Stdout() io.ReadCloser
	Stdin() io.WriteCloser
	Start() error
	Kill() error
	GetCmd() *exec.Cmd
}

type DefaultJavaExec struct {
	cmd *exec.Cmd
}

func (j *DefaultJavaExec) Stdout() io.ReadCloser {
	r, _ := j.cmd.StdoutPipe()
	return r
}

func (j *DefaultJavaExec) Stdin() io.WriteCloser {
	w, _ := j.cmd.StdinPipe()
	return w
}

func (j *DefaultJavaExec) Start() error {
	return j.cmd.Start()
}

func (j *DefaultJavaExec) Kill() error {
	return j.cmd.Process.Kill()
}

func (j *DefaultJavaExec) GetCmd() *exec.Cmd {
	return j.cmd
}

func JavaExecCmd(serverPath string, initialHeapSize, maxHeapSize int) *DefaultJavaExec {
	flags := []string{
		"-XX:MaxRAMPercentage=95",
		"-XX:+UseG1GC",
		"-XX:+ParallelRefProcEnabled",
		"-XX:MaxGCPauseMillis=200",
		"-XX:+UnlockExperimentalVMOptions",
		"-XX:+DisableExplicitGC",
		"-XX:+AlwaysPreTouch",
		"-XX:G1NewSizePercent=30",
		"-XX:G1MaxNewSizePercent=40",
		"-XX:G1HeapRegionSize=8M",
		"-XX:G1ReservePercent=20",
		"-XX:G1HeapWastePercent=5",
		"-XX:G1MixedGCCountTarget=4",
		"-XX:InitiatingHeapOccupancyPercent=15",
		"-XX:G1MixedGCLiveThresholdPercent=90",
		"-XX:G1RSetUpdatingPauseTimePercent=5",
		"-XX:SurvivorRatio=32",
		"-XX:+PerfDisableSharedMem",
		"-XX:MaxTenuringThreshold=1",
		"-Dusing.aikars.flags=https://mcflags.emc.gs",
		"-Daikars.new.flags=true",
		fmt.Sprintf("-Xms%dM", initialHeapSize), // Set heap size
		fmt.Sprintf("-Xmx%dM", maxHeapSize),     // Set heap size
		"-jar", "server.jar", "nogui"}

	cmd := exec.Command("java", flags...)
	cmd.Dir = serverPath

	return &DefaultJavaExec{cmd: cmd}
}
