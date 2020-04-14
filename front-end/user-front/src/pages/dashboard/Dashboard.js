import React from 'react';
import { Grid, Paper, makeStyles } from '@material-ui/core';
import ListaPessoa from '../pessoa/listarPessoa';


const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
    },
    paper: {
        padding: theme.spacing(2),
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
    },
}));



export default function Dashboard() {
    const classes = useStyles();
    return (
        <>
            {/* Recent Orders */}
            <Grid item xs={12}>
                <Paper className={classes.paper}>
                    {/* <Orders /> */}
                    <ListaPessoa />
                </Paper>
            </Grid>
        </>
    );
}
