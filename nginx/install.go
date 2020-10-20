package nginx

import (
	"fmt"
	"os"
	"os/exec"
)

// MeshInstance holds the information of the instance of the mesh
type MeshInstance struct {
	InstallMode     string `json:"installmode,omitempty"`
	InstallPlatform string `json:"installplatform,omitempty"`
	InstallZone     string `json:"installzone,omitempty"`
	InstallVersion  string `json:"installversion,omitempty"`
	MgmtAddr        string `json:"mgmtaddr,omitempty"`
	Nginxaddr       string `json:"nginxaddr,omitempty"`
}

// CreateInstance installs and creates a mesh environment up and running
func (h *handler) installNginx(del bool, version string) (string, error) {
	status := "installing"

	if del {
		status = "removing"
	}

	meshinstance := &MeshInstance{
		InstallVersion: version,
	}
	err := h.config.MeshInstance(meshinstance)
	if err != nil {
		return status, ErrMeshConfig(err)
	}

	h.log.Info("Installing Nginx")
	err = meshinstance.installUsingNginxctl(del)
	if err != nil {
		h.log.Err("Nginx installation failed", ErrInstallMesh(err).Error())
		return status, ErrInstallMesh(err)
	}
	if del {
		return "removed", nil
	}

	return "deployed", nil
}

// installSampleApp installs and creates a sample bookinfo application up and running
func (h *handler) installSampleApp(name string) (string, error) {
	// Needs implementation
	return "deployed", nil
}

// installMesh installs the mesh in the cluster or the target location
func (m *MeshInstance) installUsingNginxctl(del bool) error {

	Executable, err := exec.LookPath("./scripts/nginx/installer.sh")
	if err != nil {
		return err
	}

	if del {
		Executable, err = exec.LookPath("./scripts/nginx/delete.sh")
		if err != nil {
			return err
		}
	}

	cmd := &exec.Cmd{
		Path:   Executable,
		Args:   []string{Executable},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("NGINX_VERSION=%s", m.InstallVersion),
		fmt.Sprintf("NGINX_MODE=%s", m.InstallMode),
		fmt.Sprintf("NGINX_PLATFORM=%s", m.InstallPlatform),
		fmt.Sprintf("NGINX_ZONE=%s", m.InstallZone),
	)

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (m *MeshInstance) portForward() error {
	// Needs implementation
	return nil
}
