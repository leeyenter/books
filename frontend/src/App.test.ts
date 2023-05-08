import { flushPromises, mount } from "@vue/test-utils";
import App from "./App.vue";
import { router } from "./main.ts";
import { VueQueryPlugin } from "@tanstack/vue-query";
import { expect } from "vitest";

describe("App", () => {
  const wrapper = () =>
    mount(App, {
      global: {
        plugins: [router, VueQueryPlugin],
      },
    });

  it("displays table", async () => {
    const app = wrapper();
    await flushPromises();
    expect(app.text()).toContain("A Long Obedience");
  });

  it("can navigate to new book form", async () => {
    const app = wrapper();
    expect(app.text()).toContain("New");
    const newButton = app.findAll("a").find((x) => x.text().match(/New/));
    expect(newButton).toBeTruthy();
    await newButton!.trigger("click");
    await flushPromises();
    expect(app.text()).toContain("New book");
  });
});
