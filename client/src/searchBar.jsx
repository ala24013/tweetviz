import React, { useState } from "react"
import { TextInput, Box, ActionIcon, Loader } from "@mantine/core"
import { IconSearch } from "@tabler/icons"

const DEFAULT_SEARCH = "Ian"

export default function SearchBar(props) {
    const [searchValue, setSearchValue] = useState(DEFAULT_SEARCH)

    function sendMessage(val) {
        props.sendMessage(val)
    }

    return (
        <Box>
            <TextInput
                value={searchValue}
                onChange={(event) => setSearchValue(event.currentTarget.value)}
                onKeyDown={(event) => {
                    if (event.key === 'Enter') {
                        sendMessage(searchValue)
                    }
                }}
                radius="md"
                size="md"
                error={/^[0-9a-zA-Z]{0,50}$/.test(searchValue) ? null : "ASCII characters and Numbers only!"}
                rightSection={
                    props.loading ?
                        <Loader size="xs" variant="bars" />
                    :
                        <ActionIcon variant="subtle" onClick={(event) => sendMessage(searchValue)}>
                            <IconSearch />
                        </ActionIcon>
                }
                style={{ width: "40vw" }}
            />
        </Box>
    )
}