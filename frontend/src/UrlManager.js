import { useState, useEffect } from "react";
import axios from "axios";
import UrlShortener from "./UrlShortener";
import UrlList from "./UrlList";
import { Box } from "@mui/material";

const API_URL = "http://localhost:8080";

const UrlManager = () => {
  const [urls, setUrls] = useState([]);

  const fetchUrls = async () => {
    try {
      const response = await axios.get(`${API_URL}`);
      setUrls(response.data.data); 
    } catch (error) {
      console.error("Error fetching URLs:", error);
    }
  };

  // Delete URL
  const deleteUrl = async (shortenedUrl) => {
    try {
      await axios.delete(`${API_URL}/delete/${shortenedUrl}`);
      fetchUrls();
    } catch (error) {
      console.error("Error deleting URL:", error);
    }
  };

  useEffect(() => {
    fetchUrls(); // Fetch URLs on mount
  }, []);

  return (
    <Box>
      <UrlShortener refreshUrls={fetchUrls} />
      <hr/>
      <UrlList urls={urls} handleDelete={deleteUrl} />{" "}
    </Box>
  );
};

export default UrlManager;
