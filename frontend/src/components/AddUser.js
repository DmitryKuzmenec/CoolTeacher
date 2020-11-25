import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import {
    Paper, 
    TextField, 
    Grid,
    InputLabel,
    MenuItem,
    Select,
    FormControl
} from '@material-ui/core';
import DatePicker from './DatePicker'



const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
      margin: theme.spacing(3),
      width: '40ch',
    },
  },
  formControl: {
    margin: theme.spacing(3),
    minWidth: '40ch',
  },
  userInfo: {
    margin: theme.spacing(3),
    minWidth: '40ch',
  },
  selectEmpty: {
    marginTop: theme.spacing(3),
  },
}));


function AddUser() {
    const classes = useStyles();
    const [age, setAge] = React.useState('');

    function handleChange(event) {
        setAge(event.target.value);
    }
    return(
        <Paper elevation={3} >
                <Grid container justify="space-around">
                    <form  className={classes.root} noValidate autoComplete="off">
                        <TextField id="userFname" label="Имя" />
                        <TextField id="userLname" label="Фамилия" />
                        <TextField id="userEmail" label="Электронная почта" />
                        <TextField id="userPhone" label="Телефон" />
                        <FormControl className={classes.formControl}>
                            <InputLabel id="userSex-label">Пол</InputLabel>
                            <Select
                                labelId="userSex-label"
                                id="userSex"
                                value={age}
                                onChange={handleChange}
                            >
                                <MenuItem value={"m"}>Мальчик</MenuItem>
                                <MenuItem value={"f"}>Девочка</MenuItem>
                            </Select>
                        </FormControl>
                        <DatePicker label="Дата рождения"/>
                        <FormControl className={classes.userInfo}>
                            <TextField
                                id="userInfo"
                                label="Дополнительная информация"
                                placeholder="Дополнительная информация"
                                multiline
                                rows={4}
                                variant="outlined"
                            />
                        </FormControl>
                    </form>
                </Grid>
        </Paper>
    ) 
}

export default AddUser;