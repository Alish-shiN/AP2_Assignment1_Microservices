package proxy

import (
    "bytes"
    "github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
    "io"
    "net/http"
)

func forwardRequest(c *gin.Context, baseURL string) {
    client := resty.New()

    fullURL := baseURL
    if id := c.Param("id"); id != "" {
        fullURL += "/" + id
    }

    headers := make(map[string]string)
    for k, v := range c.Request.Header {
        if len(v) > 0 {
            headers[k] = v[0]
        }
    }

    req := client.R().SetHeaders(headers)

    if c.Request.Body != nil {
        bodyBytes, _ := io.ReadAll(c.Request.Body)
        c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
        req.SetBody(bodyBytes)
    }

    var resp *resty.Response
    var err error

    switch c.Request.Method {
    case http.MethodGet:
        resp, err = req.Get(fullURL)
    case http.MethodPost:
        resp, err = req.Post(fullURL)
    case http.MethodPatch:
        resp, err = req.Patch(fullURL)
    case http.MethodDelete:
        resp, err = req.Delete(fullURL)
    default:
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
        return
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
