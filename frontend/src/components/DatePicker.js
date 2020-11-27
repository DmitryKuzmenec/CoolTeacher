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
    const onChange = props['onChange'];
    const nowDate = new Date();
    const maxDate = new Date().setFullYear(nowDate.getFullYear() -1);
    const minDate = new Date().setFullYear(nowDate.getFullYear() - 19);

    const classes = useStyles();
    const [selectedDate, setSelectedDate] = React.useState(new Date('2007-01-01'));
    const [open, setOpen] = React.useState(false);
    function handleDateChange(date) {
      if (onChange !== undefined) {
        onChange(date)
      }
      setSelectedDate(date);
      setOpen(false);
    }
    
    const newLocal = <KeyboardDatePicker
      onOpen={()=>setOpen(true)}
      open={open}
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
      maxDateMessage="слишком молод"
      minDateMessage="слишком взрослый"
      invalidDateMessage="не верная дата"
      maxDate={maxDate}
      minDate={minDate}
  />
    return(
        <MuiPickersUtilsProvider utils={DateFnsUtils}>
            {newLocal}
        </MuiPickersUtilsProvider>
    ) 
}

export default DatePicker;