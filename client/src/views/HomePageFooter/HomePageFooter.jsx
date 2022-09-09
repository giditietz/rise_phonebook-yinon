import AddIcon from "@mui/icons-material/Add";
import Button from "@mui/material/Button";
import Pagination from "@mui/material/Pagination";

import text from "../../utils/language/text.json";

import "./home-page-footer.scss";

const HomePageFooter = ({ page, setPage, onAddClick, numOfPages }) => {
  return (
    <div className="home-page-footer">
      <Pagination
        count={numOfPages}
        page={page}
        onChange={(event, value) => setPage(value)}
        color="primary"
      />
      <Button
        onClick={onAddClick}
        style={{ textTransform: "none" }}
        variant="contained"
        startIcon={<AddIcon />}
      >
        {text.addContact}
      </Button>
    </div>
  );
};

export default HomePageFooter;
