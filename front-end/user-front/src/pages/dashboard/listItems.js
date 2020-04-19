import React from 'react';
import { ListItem, ListItemIcon, ListItemText, ListSubheader,  MenuItem, Typography } from '@material-ui/core';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import PeopleIcon from '@material-ui/icons/People';
import AssignmentIcon from '@material-ui/icons/Assignment';
import { Link } from 'react-router-dom';








export const mainListItems = (
    <>
        <MenuItem key="1" component={Link} to='/'>
            <ListItemIcon>
                <DashboardIcon />
            </ListItemIcon>
            <Typography >Dashboard</Typography>
        </MenuItem>
        <MenuItem  key="2" component={Link} to='/cadastrar'>
            <ListItemIcon>
                <ShoppingCartIcon />
            </ListItemIcon>
            <Typography >Dashboard 1</Typography>
        </MenuItem>
        {/* <MenuItem >
            <ListItemIcon>
                <PeopleIcon />
            </ListItemIcon>
            <Typography >Dashboard 2</Typography>
        </MenuItem> */}
    </>
);

export const secondaryListItems = (
    <div>
        <ListSubheader inset>Saved reports</ListSubheader>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Current month" />
        </ListItem>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Last quarter" />
        </ListItem>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Year-end sale" />
        </ListItem>
    </div>
);
