import TextBox from "./Textbox";
import Title from "./Title";

const FormField = ({ title, value, onChange }) => {
  return (
    <>
      <Title title={title} />
      <TextBox value={value} onChange={onChange} />
    </>
  );
};

export default FormField;
