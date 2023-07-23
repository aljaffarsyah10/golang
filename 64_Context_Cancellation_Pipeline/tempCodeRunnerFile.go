    ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
    defer cancel()
    generateFilesWithContext(ctx)
