import {
  List,
  ListItem,
  ListItemText,
  IconButton,
  Link,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

const API_URL = "http://localhost:8080";

const UrlList = ({ urls, handleDelete }) => {
  return (
    <List>
      {urls.map((url) => (
        <ListItem key={url.id}>
          <img src={"https://www.google.com/s2/favicons?domain_url=" + url.url} alt="Domain"/>
          <Link href={`${API_URL}/${url.shortened_url}`} sx={{width: "250px", marginLeft: "10px"}}>
            {"localhost:8080/" + url.shortened_url}
          </Link>
          <ListItemText>{url.title}</ListItemText>
          <IconButton
            edge="end"
            aria-label="delete"
            onClick={() => handleDelete(url.shortened_url)} // Delete the URL when clicked
          >
            <DeleteIcon />
          </IconButton>
        </ListItem>
      ))}
    </List>
  );
};

export default UrlList;
