# craft-naming

## Goal
The goal of this project is to understand what the main method does and to rename everything to make it more understandable.


## Good practices
Here are some good practices on naming to help you.


Always remember that the next rules are only guidelines and should not be taken as absolute.
Naming depends on the language and conventions you may find in the community and you team.
In doubt, always refer to your peers more than this file and write down what you decide.
Consistency is also important and you should try to use the same rules throughout your whole projects.

---

A good naming answers the question "what" and not "how". You want to be as clear as possible and express your intent.
Example for "how": `clientsMap`
Prefer something like: `clientsInfo`

An explicit name allows you to understand what will happen without having to read all the code.
Code that communicates clearly its intents doesn't require comments.

---

Try to find names that can be pronounced easily:
`ez2rd` could be renamed `easyToRead`

You can use abbreviations if everyone is supposed to know them:
`numberOfClients` could be renamed as `clientsNb`

---

As far as length goes, names should be as short as possible and as long as required.

Length can also vary from context.
A local scope calls for less ambiguity and smaller names while a more global scope requires more precise names.
`for client in clientsToSave`

From another point of view, a public function should be rather concise while a private function might be so precise it
requires a very long name.
`saveClient` vs `replaceMissingFieldsWithDefaults`

---

Names that can be searched make you save time so don't always use the same names for everything.
Always use explicit and up to date names. Remember that things can (and probably will) change and you want to be clear 
when communicating with your client or PO.

---

Don't give numbers, try to be more explicit.
`c1` -> `requestedClient`

---

Conventions by type:

* class/interface - use nouns: `Client`, `ClientValidator`
* instance variable - use nouns: `clientAddress`, `clientName`
* boolean - use state verbs: `isValid`, `hasClients`
* function - use imperative verbs: `acceptClient`, `saveCart`
* enum - use adjectives or nouns: `ACCEPTED`, `BLUE`
* event - use perfect-tense verbs: `clientSaved`, `dayChanged`




