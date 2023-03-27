import { AppBar, IconButton, Toolbar, Typography } from "@material-ui/core";
import DriverIcon from "@material-ui/icons/DriveEta";
import { FunctionComponent } from "react";

export const Navbar: FunctionComponent = () => {
  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu">
          <DriverIcon />
          <Typography variant="h6">Code Delivery</Typography>
        </IconButton>
      </Toolbar>
    </AppBar>
  );
};
