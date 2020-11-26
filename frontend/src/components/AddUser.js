import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import {
    Paper, 
    TextField, 
    Grid,
    InputLabel,
    MenuItem,
    Select,
    FormControl,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    Table,
    TableRow,
    TableCell,
    TableBody,
    TableContainer
} from '@material-ui/core';
import {Save} from '@material-ui/icons'
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
    maxWidth: '100ch',
  },
  userInfo: {
    margin: theme.spacing(3),
    minWidth: '40ch',
  },
  selectEmpty: {
    marginTop: theme.spacing(3),
  },
  button: {
      marginRight: theme.spacing(3),
      marginLeft: theme.spacing(3),
      marginBottom: theme.spacing(3),
      width: '37ch',
      height: '5ch',
  },
  table: {
    minWidth: 100,
  },
  boldText: {
    fontWeight: "bold",
  },
}));


function AddUser() {
    const classes = useStyles();
    const [sex, setSex] = React.useState('');
    const [fname, setFname] = React.useState('');
    const [lname, setLname] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [phone, setPhone] = React.useState('');
    const [birthDate, setBirthDate] = React.useState('');
    const [info, setInfo] = React.useState('');
    const [openConfirmationDialog, setOpenConfirmationDialog] = React.useState(false);
    
    function handleSexChange(e) {
        setSex(e.target.value);
    }
    function handleFnameChange(e) {
        setFname(e.target.value);
    }
    function handleLnameChange(e) {
        setLname(e.target.value);
    }
    function handleEmailChange(e) {
        setEmail(e.target.value);
    }
    function handlePhoneChange(e) {
        setPhone(e.target.value);
    }
    function handleBirthDateChange(e) {
        setBirthDate(e.target.value);
    }
    function handleInfoChange(e) {
        setInfo(e.target.value);
    }
    function handleSubmit() {
        setOpenConfirmationDialog(true);
    }
    function closeWithoutSave() {
        setOpenConfirmationDialog(false);
        console.log("no need save yet...");
    }
    function closeWithSave() {
        //console.log(data);
        setOpenConfirmationDialog(false);
    }
    let textSex = "";
    if (sex === "m") {
        textSex = "мальчик"
    } else if (sex === "f") {
        textSex = "девочка"
    }
    const dataDialog = [
        {name: "Имя", val: fname},
        {name: "Фамилия", val: lname},
        {name: "Пол", val: textSex},
        {name: "Дата рождения", val: birthDate},
        {name: "Электронная почта", val: email},
        {name: "Номер телефона", val: phone},
    ];
    return(
        <>
        <Paper  elevation={1} >
            <Grid className={classes.root} container justify="space-around">
                <TextField id="userFname" label="Имя" onChange={handleFnameChange}/>
                <TextField id="userLname" label="Фамилия" onChange={handleLnameChange}/>
                <TextField id="userEmail" label="Электронная почта" onChange={handleEmailChange}/>
                <TextField id="userPhone" label="Телефон" onChange={handlePhoneChange} />
                <FormControl className={classes.formControl}>
                    <InputLabel id="userSex-label">Пол</InputLabel>
                    <Select
                        labelId="userSex-label"
                        id="userSex"
                        value={sex}
                        onChange={handleSexChange}
                    >
                        <MenuItem value={"m"}>Мальчик</MenuItem>
                        <MenuItem value={"f"}>Девочка</MenuItem>
                    </Select>
                </FormControl>
                <DatePicker label="Дата рождения" onChange={handleBirthDateChange}/>
                <FormControl className={classes.userInfo}>
                    <TextField
                        id="userInfo"
                        label="Дополнительная информация"
                        placeholder="Дополнительная информация"
                        multiline
                        rows={4}
                        variant="outlined"
                        onChange={handleInfoChange}
                    />
                </FormControl>
                <Button
                    variant="contained"
                    color="primary"
                    size="large"
                    className={classes.button}
                    startIcon={<Save />}
                    onClick={handleSubmit}
                >
                    Save
                </Button>
            </Grid>
        </Paper>
        <Dialog
            open={openConfirmationDialog}
            //onClose={handleClose}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
        >
            <DialogTitle id="alert-dialog-title">{"Пожалуйста проверьте данные перед сохранением:"}</DialogTitle>
            <DialogContent>
            <DialogContentText id="alert-dialog-description">
                <TableContainer className={classes.table}>
                    <Table>
                        <TableBody>
                            {dataDialog.map((row) => (
                                <TableRow key={row.name}>
                                    <TableCell component="th" scope="row">
                                        <span className={classes.boldText}>{row.name}:</span>
                                    </TableCell>
                                    <TableCell align="left">{row.val}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>    
            </DialogContentText>
            </DialogContent>
            <DialogActions>
            <Button onClick={closeWithoutSave} color="primary">
                Отмена
            </Button>
            <Button onClick={closeWithSave} color="primary" autoFocus>
                Сохранить
            </Button>
            </DialogActions>
        </Dialog>
      </>
    ) 
}

export default AddUser;