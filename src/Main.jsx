import { useSelector } from 'react-redux';

import Container from '@mui/material/Container';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';

import AppBar from './features/appbar/AppBar';
import ConfigManager from './features/config/ConfigManager';
import ErrorManager from './features/errors/ErrorManager';
import ErrorSnackbar from './features/errors/ErrorSnackbar';
import ParamSetter from './features/config/ParamSetter';
import Preview from './features/preview/Preview';
import Schema from './features/schema/Schema';
import WatcherManager from './features/watcher/WatcherManager';
import Controls from './features/controls/Controls';


export default function Main() {
    const schema = useSelector(state => state.schema);

    let size = 12;
    if (schema.value.schema.length > 0) {
        size = 8;
    }

    return (
        <ErrorSnackbar >
            <WatcherManager />
            <ParamSetter />
            <ConfigManager />
            <ErrorManager />

            <AppBar />
            <Container maxWidth='xl' sx={{ marginTop: '32px' }}>
                <Box sx={{ flexGrow: 1 }}>
                    <Grid container spacing={4}>
                        <Grid size={{ xs: 12, lg: size }}>
                            <Preview scale={10} />
                            <Controls />
                        </Grid>
                        <Grid size={{ xs: 12, lg: 4 }}>
                            <Schema />
                        </Grid>
                    </Grid>
                </Box>
            </Container>
        </ErrorSnackbar>
    )
}