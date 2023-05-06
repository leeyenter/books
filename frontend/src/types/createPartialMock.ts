export const createPartialMock = <T>(fakerCreator: () => T, overrides: Partial<T>): T => {
return {...fakerCreator(), ...overrides};
}
