import React, { useState } from "react";
import axios from "axios";
import { Box, TextField, Button } from "@mui/material";

const API_URL = "http://localhost:8080";

export default function UrlShortener({ refreshUrls }) {
  const [url, setUrl] = useState("");
  const [error, setError] = useState("");

  const createUrl = async () => {
    setError(""); // Reset error before request
    if (!url.trim()) {
      setError("Please enter a valid URL.");
      return;
    }

    try {
      const response = await axios.post(`${API_URL}/shorten`, { url });
      console.log("URL created:", response.data);

      setUrl(""); // Clear input after success
      if (typeof refreshUrls === "function") {
        refreshUrls(); // ✅ Call refreshUrls only if it's a function
      } else {
        console.error("refreshUrls is not a function:", refreshUrls);
      }
    } catch (err) {
      console.error("Failed to create shortened URL:", err);
      setError("Failed to create shortened URL. Please try again.");
    }
  };

  return (
    <Box
      component="form"
      sx={{
        display: "grid",
        gridTemplateColumns: "80% auto",
        gap: "10px",
        alignItems: "center",
      }}
    >
      <TextField
        label="Paste URL here"
        variant="outlined"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <Button variant="contained" onClick={createUrl} sx={{height:"100%"}} >
        Go
      </Button>
      {error && (
        <p style={{ color: "red", textAlign: "center", marginTop: "10px" }}>
          {error}
        </p>
      )}
    </Box>
  );
}
