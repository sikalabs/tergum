package backup_process_utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sikalabs/tergum/backup_output"
	"github.com/sikalabs/tergum/backup_process"
)

func BackupProcessExecToFile(bin string, args ...string) (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return bo, err
	}
	err = bp.ExecWait(
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		bo.Stderr = stderr
		return bo, err
	}
	return bp.GetDataStderr()
}

func BackupProcessExecEnvToFile(env []string, bin string, args ...string) (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return bo, err
	}
	err = bp.ExecEnvWait(
		env,
		bin,
		args...,
	)
	if err != nil {
		stderr, _ := bp.GetStderr()
		bo.Stderr = stderr
		return bo, err
	}
	return bp.GetDataStderr()
}

func BackupProcessHttpGetWithToken(
	url string,
	tokenHeaderName string,
	tokenHeaderValue string,
	headers map[string]string,

) (backup_output.BackupOutput, error) {
	var err error
	var bo backup_output.BackupOutput

	bp := backup_process.BackupProcess{}
	bp.Init()
	err = bp.InitDataTempFile()
	if err != nil {
		return bo, err
	}

	body, err := httpGetWithToken(url, tokenHeaderName, tokenHeaderValue, headers)
	if err != nil {
		io.Copy(bp.StderrBuff, body)
		bo, _ = bp.GetDataStderr()
		return bo, err
	}

	io.Copy(bp.Data, body)
	return bp.GetDataStderr()
}

func httpGetWithToken(
	url string,
	tokenHeaderName string,
	tokenHeaderValue string,
	headers map[string]string,
) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if tokenHeaderName != "" && tokenHeaderValue != "" {
		req.Header.Set(tokenHeaderName, tokenHeaderValue)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return resp.Body, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	return resp.Body, nil
}
