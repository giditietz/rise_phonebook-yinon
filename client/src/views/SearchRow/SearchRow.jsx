import FormField from "../../components/FormField";
import text from "../../utils/language/text.json";
import "./search-row.scss";

const SearchRow = ({ firstName, setFirstName, lastName, setLastName }) => {
  return (
    <div className="search-container">
      <div className="search-field">
        <FormField
          title={text.firstName}
          value={firstName}
          onChange={setFirstName}
        />
      </div>
      <div className="search-field">
        <FormField
          title={text.lastName}
          value={lastName}
          onChange={setLastName}
        />
      </div>
    </div>
  );
};

export default SearchRow;
