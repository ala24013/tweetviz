import React, { useState } from "react"
import { TextInput, Box, ActionIcon } from "@mantine/core"
import { IconSearch } from "@tabler/icons"

const DEFAULT_SEARCH = "Ian"

export default function SearchBar(props) {
    const [searchValue, setSearchValue] = useState(DEFAULT_SEARCH)

    function sing(val) {
        console.log(val)
        console.log(props.sendMessage)
        //props.sendMessage(val)
    }

    return (
        <Box>
            <TextInput
                value={searchValue}
                onChange={(event) => setSearchValue(event.currentTarget.value)}
                onKeyDown={(event) => sing(searchValue)}
                radius="md"
                size="md"
                error={/^[0-9a-zA-Z]{0,50}$/.test(searchValue) ? null : "ASCII characters and Numbers only!"}
                rightSection={
                    <ActionIcon variant="subtle" onClick={(event) => sing(searchValue)}>
                        <IconSearch />
                    </ActionIcon>
                }
                style={{ width: "40vw" }}
            />
        </Box>
    )
}