package pkg

func JSONResponse(code int, message interface{}) map[string]interface{} {
    return map[string]interface{}{
        "status":  code,
        "message": message,
    }
}