<template>
	<div id="Home" class="columns">
		<div class="column is-4 is-offset-4">
			<b-field>
				<b-input placeholder="Search Videos"
					type="search"
					v-model="keyword"
					@keyup.native.enter="search"
					expanded>
				</b-input>
				<p class="control">
					<button class="button is-primary"
						@click="search">
						<font-awesome-icon icon="search" />
					</button>
				</p>
			</b-field>
		</div>
	</div>
</template>


<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';

const namespace: string = 'search';

@Component
export default class Home extends Vue {
	@Action('purgeData', { namespace }) public purgeData: any;
	@Action('searchVideos', { namespace }) public searchVideos: Promise<void>;

	public keyword: string = "";

	public search() {
		if (this.keyword === "") {
			return
		}

		this.purgeData();

		// @ts-ignore
		this.searchVideos(this.keyword);
	}
}
</script>
