import React from "react";
import { Box, CssBaseline, Typography } from "@mui/material";
import UrlManager from "./UrlManager";

function App() {
  return (
    <Box maxWidth="md" margin="50px auto" boxShadow="rgba(100, 100, 111, 0.2) 0px 7px 29px 0px" borderRadius="6px" padding="20px">
      <CssBaseline />
      <Typography variant="h4" align="center" gutterBottom>
        URL Shortener
      </Typography>
      <UrlManager />
    </Box>
  );
}

export default App;
