import TextField from "@mui/material/TextField";

const TextBox = ({ value, onChange }) => {
  return (
    <TextField
      value={value}
      onChange={(e) => onChange(e.target.value)}
      id="standard-basic"
      label="Standard"
      variant="standard"
    />
  );
};

export default TextBox;
