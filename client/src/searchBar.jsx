import React, { useState } from "react"
import { TextInput, Box, ActionIcon } from "@mantine/core"
import { IconSearch } from "@tabler/icons"

const DEFAULT_SEARCH = "Ian"

export default function SearchBar(props) {
    const [searchValue, setSearchValue] = useState(DEFAULT_SEARCH)

    return (
        <Box>
            <TextInput
                value={searchValue}
                onChange={(event) => setSearchValue(event.currentTarget.value)}
                radius="md"
                size="md"
                error={/^[0-9a-zA-Z]{0,50}$/.test(searchValue) ? null : "ASCII characters and Numbers only!"}
                rightSection={
                    <ActionIcon variant="subtle">
                        <IconSearch />
                    </ActionIcon>
                }
                style={{paddingTop: "10vh", width: "50vw" }}
            />
        </Box>
    )
}