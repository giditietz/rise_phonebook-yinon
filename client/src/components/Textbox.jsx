import TextField from "@mui/material/TextField";

const TextBox = ({ label, value, onChange }) => {
  return (
    <TextField
      value={value}
      onChange={(e) => onChange(e.target.value)}
      id="standard-basic"
      label={label}
      variant="standard"
    />
  );
};

export default TextBox;
