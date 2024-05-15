const useState = (init) => {
  let local = init;
  return [() => local, (value) => (local = value)];
};
