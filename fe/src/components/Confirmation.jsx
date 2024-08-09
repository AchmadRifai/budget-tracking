import { Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle } from "@mui/material"
import { useState } from "react"

export default function Confirmation({ onClose, buttonText, onYes }) {
    const [open, setOpen] = useState(false)
    const closing = () => {
        setOpen(false)
        onClose()
    }
    const agree = () => {
        setOpen(false)
        onYes()
    }
    const openning = () => setOpen(true)
    return <>
        <Button onClick={() => openning()} aria-hidden='false' variant="outlined">{buttonText}</Button>
        <Dialog onClose={() => closing()} open={open}>
            <DialogTitle>Confirmation</DialogTitle>
            <DialogContent>
                <DialogContentText>Are you sure? If you agree, We cannot undo this process</DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => agree()} color="error">Yes</Button>
                <Button onClick={() => closing()}>No</Button>
            </DialogActions>
        </Dialog>
    </>
}