import 'date-fns';
import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import DateFnsUtils from '@date-io/date-fns';
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from '@material-ui/pickers';


const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      margin: theme.spacing(3),
      width: '40ch',
    },
  },
}));


function DatePicker(props) {
    const label = props['label'];
    const classes = useStyles();
    const [selectedDate, setSelectedDate] = React.useState(new Date('2014-08-18T21:11:54'));
    function handleDateChange(date) {
        setSelectedDate(date);
    }
    
    const newLocal = <KeyboardDatePicker
    disableToolbar
    variant="inline"
    format="MM/dd/yyyy"
    margin="normal"
    id="date-picker-inline"
    label={label}
    value={selectedDate}
    onChange={handleDateChange}
    KeyboardButtonProps={{
      'aria-label': 'change date',
    }}
  />
    return(
        <MuiPickersUtilsProvider utils={DateFnsUtils}>
            {newLocal}
        </MuiPickersUtilsProvider>
    ) 
}

export default DatePicker;