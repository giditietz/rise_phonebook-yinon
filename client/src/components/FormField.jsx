import TextBox from "./Textbox";
import Title from "./Title";

const FormField = ({ title, value, onChange }) => {
  return (
    <div>
      <Title title={title} />
      <TextBox value={value} onChange={onChange} />
    </div>
  );
};

export default FormField;
