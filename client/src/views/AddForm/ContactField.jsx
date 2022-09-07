import FormField from "../../components/FormField";
import text from "../../utils/language/text.json";

const ContactFields = ({ firstName, setFirstName, lastName, SetLastName }) => {
  return (
    <>
      <FormField
        title={text.firstName}
        value={firstName}
        onChange={setFirstName}
      />
      <FormField
        title={text.lastName}
        value={lastName}
        onChange={SetLastName}
      />
    </>
  );
};

export default ContactFields;
